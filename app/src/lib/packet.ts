export class Packet {
    type: string
    action: string
    payload: string
    from: string
    to: string

    constructor(type: string, action: string, payload: string, from: string, to: string) {
        this.type = type
        this.action = action
        this.payload = payload
        this.from = from
        this.to = to
    }
}