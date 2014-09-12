require.config({
	baseUrl: '',
	paths: {
		//REQUIRE LIBRARIES
		'jquery': '/js/bower_components/jquery/dist/jquery.min',
		'react': '/js/bower_components/react/react'

	},
});

require(['/js/components/views/index.js'], function (init) { init() })