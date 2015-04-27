'use strict';

/**
 * @ngdoc overview
 * @name clientApp
 * @description
 * # clientApp
 *
 * Main module of the application.
 */
angular
    .module('clientApp', [
        'ngAnimate',
        'ngCookies',
        'ngResource',
        'ngRoute',
        'ngSanitize',
        'ngTouch',
        'ui.grid'
    ])
    .constant('config', {
        url: 'http://localhost:8080'
    })
    .factory('myInterceptor', ['$rootScope', '$q', '$timeout', '$cookies', function ($rootScope, $q, $timeout, $cookies) {
        var myInterceptor = {
            'request': function (config) {
                var token = $cookies['XSRF-TOKEN'];
                if (token) {
                    config.headers['X-XSRF-TOKEN'] = token;
                    config.headers['X-Xsrftoken'] = token;
                }
                config.headers['X-Xsrftoken'] = 'ddddd';
                return config;
            },
            'responseError': function (response) {
                console.log('status ' + response.status);

                if (response.status === 401) {
                    $rootScope.$broadcast('InvalidToken');
                    $rootScope.sessionExpired = true;
                    $timeout(function () {
                        $rootScope.sessionExpired = false;
                    }, 5000);
                } else if (response.status === 403) {
                    $rootScope.$broadcast('InsufficientPrivileges');
                } else {
                    // Here you could handle other status codes, e.g. retry a 404
                }
                // return response;
                return $q.reject(response);
            }
        };
        return myInterceptor;
    }])
    .config(['$httpProvider', function($httpProvider) {
        $httpProvider.interceptors.push('myInterceptor');
    }])
    .config(['$routeProvider', function ($routeProvider) {
        $routeProvider
            .when('/', {
                templateUrl: 'views/main.html',
                controller: 'MainCtrl'
            })
            .when('/login', {
                templateUrl: 'views/login.html',
                controller: 'LoginCtrl'
            })
            .otherwise({
                redirectTo: '/'
            });
    }]);
