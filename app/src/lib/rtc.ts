import { addMessageToChat } from '$lib/index.ts'
import { get } from 'svelte/store'
import { Packet } from './packet'
import { connectionStatus, selectedPeer, username } from './stores'
import { sendPacket } from './socket'

export class RTC {
	conn: RTCPeerConnection
	chan: RTCDataChannel
	msgQueue: string[]
	peerName: string

	constructor(peerName: string) {
		this.peerName = peerName
		this.msgQueue = []

		this.conn = new RTCPeerConnection({
			// udp hole punching
			iceServers: [
				{ urls: 'stun:stun.l.google.com:19302' },
			],
		})

		this.conn.ondatachannel = (e) => {
			this.chan = e.channel
		}

		this.chan = this.conn.createDataChannel(peerName)
		
		this.chan.onmessage = (e) => {
			this.handleIncomingMessage(e.data)
		}

		this.chan.onopen = () => {
			this.handleConnectionOpened()
		}

		this.chan.onclose = () => {
			this.handleConnectionClosed()
		}
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
		await this.conn.setRemoteDescription(JSON.parse(answer))
	}

	sendMessage(message: string) {
		if (this.chan.readyState != 'open') {
			console.log('adding to queue')
			this.msgQueue.push(message)
		}
		else {
			this.chan.send(message)
		}
	}

	sendQueuedMessages() {
		console.log('sending queued messages: ', this.msgQueue)
		while (this.msgQueue.length > 0) {
			let msg = this.msgQueue.pop()
			this.chan.send(msg || '')
		}
	}

	handleConnectionOpened() {
		console.log('connection opened with:', this.peerName)
		connectionStatus.set('open')
		this.sendQueuedMessages()
	}

	handleIncomingMessage(message: string) {
		addMessageToChat(message, this.peerName, this.peerName)
	}

	handleConnectionClosed() {
		console.log('answering channel closed.')
		connectionStatus.set('closed')
	}

	async sendOffer() {
		sendPacket(new Packet('passPacket', await this.makeOffer(), get(username), get(selectedPeer).name));
		console.log('offer sent to:', this.peerName);
	}

	async sendAnswer(offer: string) {
		sendPacket(new Packet('passPacket', await this.makeAnswer(offer), get(username) || '', get(selectedPeer).name))
		console.log('answer sent to:', this.peerName)
	}
}
