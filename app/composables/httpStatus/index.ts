export const useHttp = () => {
    return useState('http', () => code)
}
const code = {
    InternalServerError:500,
    BadRequest:400,
}