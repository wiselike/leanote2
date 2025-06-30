/* This script make #search-result-wrapper switch to unloaded or shown automatically. */

$(function() {

	const btnSbTrigger = $("#sidebar-trigger");
	const btnSearchTrigger = $("#search-trigger");
	const btnCancel = $("#search-cancel");
	const btnClear = $("#search-cleaner");

	const topbarTitle = $("#topbar-title");
	const searchWrapper = $("#search-wrapper");
	const input = $("#search-input");
	const hints = $("#search-hints");


	/*--- Actions in small screens (Sidebar unloaded) ---*/

	const mobileSearchBar = (function () {
		return {
			on() {
				btnSbTrigger.addClass("unloaded");
				topbarTitle.addClass("unloaded");
				btnSearchTrigger.addClass("unloaded");
				searchWrapper.addClass("d-flex");
				btnCancel.addClass("loaded");
			},
			off() {
				btnCancel.removeClass("loaded");
				searchWrapper.removeClass("d-flex");
				btnSbTrigger.removeClass("unloaded");
				topbarTitle.removeClass("unloaded");
				btnSearchTrigger.removeClass("unloaded");
			}
		};
	}());

	function isMobileView() {
		return btnCancel.hasClass("loaded");
	}

	btnSearchTrigger.click(function() {
		mobileSearchBar.on();
		input.focus();
	});

	btnCancel.click(function() {
		mobileSearchBar.off();
	});

	input.focus(function() {
		searchWrapper.addClass("input-focus");
	});

	input.focusout(function() {
		searchWrapper.removeClass("input-focus");
	});

	input.on("keyup", function(e) {
		if (e.keyCode === 8 && input.val() === "") {
			if (!isMobileView()) {
			} else {
				hints.removeClass("unloaded");
			}
		} else {
			if (input.val() !== "") {

				if (!btnClear.hasClass("visible")) {
					btnClear.addClass("visible");
				}

				if (isMobileView()) {
					hints.addClass("unloaded");
				}
			}
		}
	});

	btnClear.on("click", function() {
		input.val("");
		if (isMobileView()) {
			hints.removeClass("unloaded");
		} else {
		}
		input.focus();
		btnClear.removeClass("visible");
	});

});
