<script lang='ts'>
    import { sendPacket } from '$lib/socket.ts'
    import { username } from '$lib/stores.ts'
    import { onMount } from 'svelte'
    import { Packet } from '$lib/packet.ts'

    let usernameInput: string = ''
    let usernameValidity: string = 'Valid'

    export let waitForRegistration: Function

    onMount( () => {
		// Check if the username for registration is valid
        const usernameInputElement = document.getElementById('username-input') as HTMLInputElement
        usernameInputElement.addEventListener('input', () => {
            checkUsernameValidity()
        })
    })

    function checkUsernameValidity(): void {
        if (usernameInput.length < 3) {
            usernameValidity = 'Username too short.'
        } else if (/[^a-zA-Z0-9]/.test(usernameInput)) {
            usernameValidity = 'Username may only contain letters and numbers.'
        } else {
            usernameValidity = 'Valid'
        }
    }

    function register(): void {
		checkUsernameValidity()
		if (usernameValidity == 'Valid') {
			waitForRegistration(true)
			username.set(usernameInput)
			localStorage.setItem('username', usernameInput)
            let registrationPacket = new Packet(
                'post',
                'register',
                '',
                usernameInput,
                'server'
            )
            console.log('sent registration packet')
            sendPacket(registrationPacket)
		}
    }

</script>

<form
    on:submit|preventDefault={() => register()}
    class="w-1/3 place-self-center card p-4 text-token space-y-4 flex flex-col"
>
    <h2 class="h2 pb-2 font-bold">Register</h2>
    <input
        bind:value={usernameInput}
        class="input"
        aria-autocomplete="none"
        type="text"
        placeholder="Username"
        id="username-input"
        maxlength=16
        autocomplete="off"
        autocorrect="off"
        autocapitalize="off"
        spellcheck="false"
    />

    {#if usernameValidity != "Valid"}
        <p class="text-red-500 mt-2">{usernameValidity}</p>
    {/if}

    <button type="submit" class="w-full btn variant-filled text-center">Register</button>

    <div>
        <span class="pt-10">Have an account already?</span>
        <a href="../login" class="text-blue-500">
            Log in
        </a>
    </div>
</form>