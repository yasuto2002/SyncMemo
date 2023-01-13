package auth

import (
	"context"
	_ "embed"
	"fmt"
	"net/http"
	"syncmemo/clock"
	"syncmemo/entity"
	"syncmemo/store"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

//go:embed cert/public.pem
var pubKey []byte

//go:embed cert/secret.pem
var privateKey []byte

const (
	UserMailKey = "user_mail"
)

type Jwts struct {
	PubKey string
}

type JWTer struct {
	PrivateKey, PublicKey jwk.Key
	Clocker               clock.Clocker
}

func (j *JWTer) GenerateToken(ctx context.Context, kvs *redis.Client, u entity.User) ([]byte, error) {
	tok, err := jwt.NewBuilder().JwtID(uuid.NewString()).Subject("access_token").IssuedAt(j.Clocker.Now()).Expiration(j.Clocker.Now().Add(30*time.Minute)).Claim(UserMailKey, u.Mail).Build()
	if err != nil {
		return nil, fmt.Errorf("GenerateToken: failed to build token: %w", err)
	}
	if err := store.Save(ctx, kvs, tok.JwtID(), u.Mail); err != nil {
		return nil, err
	}
	signed, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, j.PrivateKey))
	if err != nil {
		return nil, err
	}
	return signed, nil
}

func parse(rawKey []byte) (jwk.Key, error) {
	key, err := jwk.ParseKey(rawKey, jwk.WithPEM(true))
	if err != nil {
		return nil, err
	}
	return key, nil
}

func NewJWTer(c clock.Clocker, kvs *redis.Client) (*JWTer, error) {
	j := &JWTer{}
	privkey, err := parse(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: private key: %w", err)
	}
	pubkey, err := parse(pubKey)

	if err != nil {
		return nil, fmt.Errorf("failed in NewJWTer: public key: %w", err)
	}

	j.PrivateKey = privkey
	j.PublicKey = pubkey
	j.Clocker = c
	return j, nil

}

func (j *JWTer) GetToken(ctx context.Context, rdb *redis.Client, r *http.Request) (jwt.Token, error) {
	//httpリクエストからトークンを取り出す
	token, err := jwt.ParseRequest(
		r,
		jwt.WithKey(jwa.RS256, j.PublicKey),
		jwt.WithValidate(false),
	)
	if err != nil {
		return nil, err
	}
	if err := jwt.Validate(token, jwt.WithClock(j.Clocker)); err != nil {
		return nil, fmt.Errorf("GetToken: failed to validate token: %w", err)
	}

	if _, err := store.Load(ctx, rdb, token.JwtID()); err != nil {
		return nil, fmt.Errorf("GetToken: %q expired: %w", token.JwtID(), err)
	}
	return token, nil
}

type userMailKey struct{}

func (j *JWTer) FillContext(r *http.Request, rdb *redis.Client) (*http.Request, error) {
	token, err := j.GetToken(r.Context(), rdb, r)
	if err != nil {
		return nil, err
	}
	uid, err := store.Load(r.Context(), rdb, token.JwtID())
	if err != nil {
		return nil, err
	}
	ctx := SetUserID(r.Context(), uid)
	clone := r.Clone(ctx)
	return clone, nil
}

func SetUserID(ctx context.Context, uid string) context.Context {
	return context.WithValue(ctx, userMailKey{}, uid)
}

func GetUserID(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(userMailKey{}).(string)
	return id, ok
}
