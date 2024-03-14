export class Packet {
    action: string
    payload: string
    from: string
    to: string

    constructor(action: string, payload: string, from: string, to: string) {
        this.action = action
        this.payload = payload
        this.from = from
        this.to = to
    }
}