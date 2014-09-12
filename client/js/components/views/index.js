/** @jsx React.DOM */ 
define([
'jquery',
'react',
'js/components/views/container.js'
],

function ($, React, Container) {

	function init () {
		React.renderComponent(<Container />, $('body')[0]);
	}

	return init;
})