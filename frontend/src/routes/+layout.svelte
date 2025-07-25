<script lang="ts">
	import '../app.css';
	import { afterNavigate } from '$app/navigation';
	import { page } from '$app/state';
	import NavLink from '$lib/NavLink.svelte';
	import { Check, Ban, BellOff, Phone, UtensilsCrossed, Cigarette, Server } from '@lucide/svelte';

	let { children } = $props();

	const navItems = [
		{ Name: 'Start', Link: '/' },
		{ Name: 'Einkauf', Link: '/Einkauf' },
		{ Name: 'Mitarbeiter', Link: '/Mitarbeiter' },
		{ Name: 'Lieferanten', Link: '/Lieferanten' },
		{ Name: 'Formulare', Link: '/Formular' },
		{ Name: 'CE Archiv', Link: '/Archiv' },
		{ Name: 'Kunden', Link: '/Kunden' },
		{ Name: 'Warenlieferung', Link: '/Warenlieferung' },
		{ Name: 'CMS', Link: '/CMS' },
		{ Name: 'SN', Link: '/Seriennummer' },
		{ Name: 'Info', Link: '/Info' },
		{ Name: 'Label', Link: '/Label' },
		{ Name: 'Aussteller', Link: '/Aussteller' }
	];

	let uri = $state('');
	let status = $state('');

	$effect(() => {
		uri = page.url.pathname;
	});

	const getStatus = async () => {
		const url = 'https://status.computer-extra.net/status.php';
		const res = await fetch(url);
		const json = await res.json();
		const data = json[0];
		return data.status as string;
	};

	$effect(() => {
		const sync = setInterval(async () => {
			const res = await getStatus();
			status = res;
		}, 5 * 1000);

		return () => {
			clearInterval(sync);
		};
	});

	afterNavigate(() => {
		window.HSStaticMethods.autoInit();
	});
</script>

<header
	class="relative flex w-full flex-wrap bg-white py-3 text-sm sm:flex-nowrap sm:justify-start dark:bg-neutral-800 print:hidden"
>
	<nav class="container mx-auto px-4 sm:flex sm:items-center sm:justify-between">
		<div class="flex items-center justify-between">
			<a
				class="focus:outline-hidden flex-none text-xl font-semibold focus:opacity-80 dark:text-white"
				href="/"
				aria-label="Brand"
			>
				Viktor
			</a>
			<div class="sm:hidden">
				<button
					type="button"
					class="hs-collapse-toggle shadow-2xs focus:outline-hidden relative flex size-9 items-center justify-center gap-x-2 rounded-lg border border-gray-200 bg-white text-gray-800 hover:bg-gray-50 focus:bg-gray-50 disabled:pointer-events-none disabled:opacity-50 dark:border-neutral-700 dark:bg-transparent dark:text-white dark:hover:bg-white/10 dark:focus:bg-white/10"
					id="hs-navbar-example-collapse"
					aria-expanded="false"
					aria-controls="hs-navbar-example"
					aria-label="Toggle navigation"
					data-hs-collapse="#hs-navbar-example"
				>
					<svg
						class="hs-collapse-open:hidden size-4 shrink-0"
						xmlns="http://www.w3.org/2000/svg"
						width="24"
						height="24"
						viewBox="0 0 24 24"
						fill="none"
						stroke="currentColor"
						stroke-width="2"
						stroke-linecap="round"
						stroke-linejoin="round"
						><line x1="3" x2="21" y1="6" y2="6"></line><line x1="3" x2="21" y1="12" y2="12"
						></line><line x1="3" x2="21" y1="18" y2="18"></line></svg
					>
					<svg
						class="hs-collapse-open:block hidden size-4 shrink-0"
						xmlns="http://www.w3.org/2000/svg"
						width="24"
						height="24"
						viewBox="0 0 24 24"
						fill="none"
						stroke="currentColor"
						stroke-width="2"
						stroke-linecap="round"
						stroke-linejoin="round"><path d="M18 6 6 18"></path><path d="m6 6 12 12"></path></svg
					>
					<span class="sr-only">Toggle navigation</span>
				</button>
			</div>
		</div>
		<div
			id="hs-navbar-example"
			class="hs-collapse hidden grow basis-full overflow-hidden transition-all duration-300 sm:block"
			aria-labelledby="hs-navbar-example-collapse"
		>
			<div
				class="mt-5 flex flex-col gap-5 sm:mt-0 sm:flex-row sm:items-center sm:justify-end sm:ps-5"
			>
				{#each navItems as item}
					<NavLink href={item.Link} title={item.Name} path={uri} />
				{/each}
			</div>
		</div>
	</nav>
</header>
<div class="container mx-auto mt-5">
	{@render children()}
</div>
<div style="position: fixed; bottom: 1rem; right: 1rem;" class="print:hidden">
	<p class="rounded border border-gray-200 p-2">
		{#if status.length > 3}
			<span class="flex flex-row items-center gap-x-2">
				Johannes Status: {status}
				{#if status === 'Anwesend'}
					<Check class="!text-green-500" />
				{:else if status === 'Abwesend'}
					<Ban class="!text-red-500" />
				{:else if status === 'Beschäftigt'}
					<BellOff class="!text-red-500" />
				{:else if status === 'Am Telefonieren'}
					<Phone class="!text-blue-500" />
				{:else if status === 'Im Mittag'}
					<UtensilsCrossed class="!text-red-500" />
				{:else if status === 'Am Rauchen'}
					<Cigarette class="!text-red-500" />
				{:else if status === 'Auf Silbers Platz'}
					<UtensilsCrossed class="!text-red-500" />
				{:else if status === 'In Wartung'}
					<Server class="!text-red-500" />
				{/if}
			</span>
		{:else}
			<span id="status">LOADING ...</span>
		{/if}
	</p>
</div>
