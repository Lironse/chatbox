import { addMessageToChat } from '$lib/index.ts'
import { selectedPeer, savedPeers } from '$lib/stores.ts'
import { Peer } from '$lib/peer.ts'
import { get } from 'svelte/store'

export class RTC {
	conn: RTCPeerConnection
	chan: RTCDataChannel
    msgQueue: string[]
	from: string

	constructor() {
		this.from = 'unset'
        this.msgQueue = []
		this.conn = new RTCPeerConnection({
			iceServers: [
				{ urls: 'stun:stun.l.google.com:19302' },
			],
		})
		this.chan = this.conn.createDataChannel('chat')
		this.chan.onmessage = (e) => this.handleIncomingMessage(e.data)
		this.chan.onopen = (e) => {
			console.log('initial channel opened')
			this.onConnection()
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
				this.onConnection()
			}
			this.chan.onclose = (e) => console.log('answering channel closed.')
		}

		await this.conn.setRemoteDescription(JSON.parse(offer))
		this.conn.createAnswer().then((answer) => this.conn.setLocalDescription(answer))

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
		let peer = get(selectedPeer)
        addMessageToChat(message, peer.name, peer.avatar)
    }

	onConnection() {
		let newPeer = new Peer(this.from, this.from)
		selectedPeer.set(newPeer)
		
		let peers = get(savedPeers)
		const peerExists = peers.some((peer) => peer.name === newPeer.name)

		if (!peerExists) {
			// Peer doesn't exist, add it to savedPeers
			savedPeers.update((peers) => [...peers, newPeer]) // Adjust update function according to your store implementation
			localStorage.setItem('savedPeers', JSON.stringify(peers))
		}
		
		this.sendQueuedMessages()
	}
}


