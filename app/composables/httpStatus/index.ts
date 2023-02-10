import { errCode } from "~~/repository/errCode"
export const useHttp = () => {
    return useState('http', () => code)
}
interface errCodes {
    Success: errCode
    InternalServerError: errCode
    BadRequest:errCode
    Unauthorized: errCode;
}
const code:errCodes = {
    Success:200,
    InternalServerError:500,
    BadRequest:400,
    Unauthorized:401
}