<script lang="ts">
	import { username, rtc } from '$lib/stores.ts'
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
		autocomplete="off"
		spellcheck="false"
	/>
	<input type="submit" hidden />
</form>
