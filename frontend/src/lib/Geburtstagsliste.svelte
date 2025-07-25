<script lang="ts">
	import { GetAllMitarbeiter } from './wailsjs/go/main/App';
	import type { main } from './wailsjs/go/models';

	const checkTime = (input: number) => (input < 10 ? '0' + input : input.toString());
	const startTime = () =>
		`${new Date().getHours()}:${checkTime(new Date().getMinutes())}:${checkTime(new Date().getSeconds())}`;

	let time = $state(startTime());
	let mitarbeiter: main.Geburtstag | undefined = $state();

	// get Mitarbeiter
	$effect(() => {
		GetAllMitarbeiter().then((res) => {
			console.log(res);
			mitarbeiter = res;
		});
	});

	// Starte Uhrzeit
	$effect(() => {
		const timeInterval = setInterval(() => {
			time = startTime();
		}, 1000);

		return () => {
			clearInterval(timeInterval);
		};
	});
</script>

<h2 class="mt-5">
	Heute ist der {new Date().toLocaleDateString('de-de', {
		day: '2-digit',
		month: '2-digit',
		year: 'numeric'
	})} - {time}
</h2>

{#if mitarbeiter?.Heute && mitarbeiter.Heute.length > 0}
	<div class="my-5">
		{#each mitarbeiter?.Heute as ma}
			<div
				class="rounded-lg border-t-2 border-teal-500 bg-teal-50 p-4 dark:bg-teal-800/30"
				role="alert"
				tabindex="-1"
				aria-labelledby="hs-bordered-success-style-label"
			>
				<div class="flex">
					<div class="shrink-0">
						<span
							class="inline-flex size-8 items-center justify-center rounded-full border-4 border-teal-100 bg-teal-200 text-teal-800 dark:border-teal-900 dark:bg-teal-800 dark:text-teal-400"
						>
							<svg
								class="size-4 shrink-0"
								xmlns="http://www.w3.org/2000/svg"
								width="24"
								height="24"
								viewBox="0 0 24 24"
								fill="none"
								stroke="currentColor"
								stroke-width="2"
								stroke-linecap="round"
								stroke-linejoin="round"
							>
								<path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2 2 6.477 2 12s4.477 10 10 10z"
								></path>
								<path d="m9 12 2 2 4-4"></path>
							</svg>
						</span>
					</div>
					<div class="ms-3">
						<h3
							id="hs-bordered-success-style-label"
							class="font-semibold text-gray-800 dark:text-white"
						>
							{ma.Name}
						</h3>
						<p class="text-sm text-gray-700 dark:text-neutral-400">hat heute Geburtstag</p>
					</div>
				</div>
			</div>
		{/each}
	</div>
{/if}

{#if mitarbeiter?.Zukunft && mitarbeiter.Zukunft.length > 0}
	<h1>Zukünftige Geburtstage</h1>
{/if}

{#if mitarbeiter?.Vergangenheit && mitarbeiter.Vergangenheit.length > 0}
	<h1>Zukünftige Geburtstage</h1>
{/if}
