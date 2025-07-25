
// See https://svelte.dev/docs/kit/types#app.d.ts
// for information about these interfaces
declare global {
	interface Window {
		_;
		$: typeof import("jquery");
    	jQuery: typeof import("jquery");
    	DataTable;
    	Dropzone;
    	VanillaCalendarPro;
		noUiSlider;

		// Preline UI
    	HSStaticMethods: IStaticMethods;
	}

	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

export { };

