export interface BoardHistory {
    id:string
    name:string
    mail:string
    password:string
    createdAt:string
}
export interface Boards {
    boards:Array<BoardHistory>
}

