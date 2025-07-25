// Optional plugins
import "datatables.net";
import "dropzone/dist/dropzone-min.js";
import $ from "jquery";
import _ from "lodash";
import noUiSlider from "nouislider";
import * as VanillaCalendarPro from "vanilla-calendar-pro";

window._ = _;
window.$ = $;
window.jQuery = $;
window.DataTable = $.fn.dataTable;
window.noUiSlider = noUiSlider;
window.VanillaCalendarPro = VanillaCalendarPro;

// Preline UI
import("preline/dist");