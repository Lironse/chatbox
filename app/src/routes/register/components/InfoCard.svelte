<script lang='ts'>
    import { username } from '$lib/stores.ts'
</script>

<div class="grid h-full">
    <div class="w-1/3 place-self-center card p-4 text-token space-y-4 flex flex-col">
        <h2 class="h2">
            <span class="bg-gradient-to-br from-red-500 to-pink-300 bg-clip-text text-transparent box-decoration-clone">Welcome, {$username}!</span>
        </h2>
        <span>your private key has been saved to some location.</span>
        <div class="flex justify-end">
            <a type="button" href="/chat" class="btn variant-filled">Continue</a>
        </div>
    </div>
</div>