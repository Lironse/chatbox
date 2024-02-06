<script lang="ts">
	import { messageFeed, username, rtc } from '$lib/stores.ts'
    import { addMessageToChat } from '$lib/index.ts'

	let addMessageInput: string
    
    function addMessage() {
        addMessageToChat(addMessageInput, $username, $username)
        rtc.sendMessage(addMessageInput)
        addMessageInput = ''
    }

</script>

<form class="p-4" name="addPeer" on:submit={addMessage}>
	<input
		id='message-input'
		bind:value={addMessageInput}
		class="input"
		type="text"
		placeholder="Enter message"
		maxlength="50"
	/>
	<input type="submit" hidden />
</form>
