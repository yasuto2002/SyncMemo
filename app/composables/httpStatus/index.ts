export const useHttp = () => {
    return useState('http', () => code)
}
const code = {
    Success:200,
    InternalServerError:500,
    BadRequest:400,
}