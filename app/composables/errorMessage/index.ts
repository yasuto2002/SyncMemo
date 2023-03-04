export const useMail = () => {
  return useState("mail", () => "有効なメールアドレスを入力してください")
}
export const useRequired = () => {
  return useState("required", () => "必須項目です")
}
export const useMax = () => {
  return useState("max", () => "最大文字数は10文字です")
}
const mes = {
  mail: "有効なメールアドレスを入力してください",
  max: "文字数の上限を超えています",
  min: "文字数の下限を下回っています",
  regex:
    /^(?!.*(\<|\>|\"|#|\$|%|&|\'|\(|\)|\*|\+|\,|\/|\:|\;|=|\[|\\|\]|_|\`|{|\||}|~|\!))/,
  regexMes: "不正な文字列が含まれています",
  required: "必須項目です",
  match: "パスワードが一致しません",
}
export const useValidateMes = () => {
  return useState("validate", () => mes)
}
