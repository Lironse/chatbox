import { addMessageToChat } from '$lib/index.ts'
import { selectedPeer } from '$lib/stores.ts'
import { Peer } from '$lib/peer.ts'


export class RTC {
	conn: RTCPeerConnection;
	chan: RTCDataChannel;
    msgQueue: string[];
	from: string;

	constructor() {
		this.from = 'unset'
        this.msgQueue = []
		this.conn = new RTCPeerConnection()
		this.chan = this.conn.createDataChannel('chat')
		this.chan.onmessage = (e) => this.handleIncomingMessage(e.data)
		this.chan.onopen = (e) => {
			console.log('initial channel opened')
			selectedPeer.set(new Peer(this.from, this.from))
			this.sendQueuedMessages()
		}
		this.chan.onclose = (e) => console.log('initial channel closed.')
	}

	async makeOffer(): Promise<string> {
		const offer = await this.conn.createOffer()
		await this.conn.setLocalDescription(offer)

		await new Promise<void>((resolve) => {
			this.conn.onicegatheringstatechange = () => {
				if (this.conn.iceGatheringState === 'complete') {
					resolve()
				}
			}
		})

		const offerSdp: string = JSON.stringify(this.conn.localDescription)
		return offerSdp
	}

	async makeAnswer(offer: string): Promise<string> {
		this.chan.close()
		this.conn.ondatachannel = (e) => {
			this.chan = e.channel
			this.chan.onmessage = (e) => this.handleIncomingMessage(e.data)
			this.chan.onopen = (e) => {
				console.log('answering channel opened')
				selectedPeer.set(new Peer(this.from, this.from))
				this.sendQueuedMessages()
			}
			this.chan.onclose = (e) => console.log('answering channel closed.')
		};

		await this.conn.setRemoteDescription(JSON.parse(offer))
		this.conn.createAnswer().then((answer) => this.conn.setLocalDescription(answer));

		await new Promise<void>((resolve) => {
			this.conn.onicegatheringstatechange = () => {
				if (this.conn.iceGatheringState === 'complete') {
					resolve()
				}
			}
		})

		const answerSdp = JSON.stringify(this.conn.localDescription)
		return answerSdp
	}

	async acceptAnswer(answer: string) {
		this.conn.setRemoteDescription(JSON.parse(answer))
	}

	async sendMessage(message: string) {
        if (this.chan.readyState != 'open') {
            this.msgQueue.push(message)
        }
        else {
            this.chan.send(message)
        }
	}

    sendQueuedMessages() {
        while (this.msgQueue.length > 0) {
            let msg = this.msgQueue.pop()
            this.chan.send(msg || '')
        }
    }

    handleIncomingMessage(message: string) {
        console.log('received:', message)
		let peer = {name: 'unset', avatar: 'unset'}
		selectedPeer.subscribe((value) => {
			peer = value
		})
        addMessageToChat(message, peer.name, peer.avatar)
    }
}


