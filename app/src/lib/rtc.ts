import { addMessageToChat } from '$lib/index.ts';

export class RTC {
	conn: RTCPeerConnection;
	chan: RTCDataChannel;
    msgQueue: string[];

	constructor() {
        this.msgQueue = [];
		this.conn = new RTCPeerConnection();
		this.chan = this.conn.createDataChannel('chat');
		this.chan.onmessage = (e) => this.handleIncomingMessage(e.data);
		this.chan.onopen = (e) => this.sendQueuedMessages();
		this.chan.onclose = (e) => console.log('channel closed.');
	}

	async makeOffer(): Promise<string> {
		const offer = await this.conn.createOffer();
		await this.conn.setLocalDescription(offer);

		await new Promise<void>((resolve) => {
			this.conn.onicegatheringstatechange = () => {
				if (this.conn.iceGatheringState === 'complete') {
					resolve();
				}
			};
		});

		const offerSdp: string = JSON.stringify(this.conn.localDescription);
		return offerSdp;
	}

	async makeAnswer(offer: string): Promise<string> {
		this.chan.close();
		this.conn.ondatachannel = (e) => {
			this.chan = e.channel;

			this.chan.onmessage = (e) => this.handleIncomingMessage(e.data);
			this.chan.onopen = (e) => this.sendQueuedMessages();
			this.chan.onclose = (e) => console.log('channel closed.');
		};

		await this.conn.setRemoteDescription(JSON.parse(offer));
		this.conn.createAnswer().then((answer) => this.conn.setLocalDescription(answer));

		await new Promise<void>((resolve) => {
			this.conn.onicegatheringstatechange = () => {
				if (this.conn.iceGatheringState === 'complete') {
					resolve();
				}
			};
		});

		const answerSdp = JSON.stringify(this.conn.localDescription);
		return answerSdp;
	}

	async acceptAnswer(answer: string) {
		this.conn.setRemoteDescription(JSON.parse(answer));
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
        addMessageToChat(message, 'peer', '2')
    }
}


