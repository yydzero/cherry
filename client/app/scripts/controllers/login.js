'use strict';


angular.module('clientApp').controller('LoginCtrl', ['$scope', '$http', 'config', '$cookies', '$location',
    function ($scope, $http, config, $cookies, $location) {
        $scope.message = '';

        $scope.login = function (email, password) {
            if (_.isEmpty(email) || _.isEmpty(password)) {
                $scope.message = 'username and password could not be empty';
                return;
            }

            $http.get(config.url + '/login?email=' + email + '&password=' + password)
                .success(function (data) {
                    console.log(data);
                    $cookies['XSRF-TOKEN'] = data['sessionid'];
                    $cookies['USER'] = data.user;

                    console.log('signin is finished with ' + data);

                    //$state.go('projects');
                    $location.path('/');

                })
                .error(function (error) {
                    $scope.message = 'failed: ' + error;
                });

        };

        $scope.logout = function() {
            $http.post(config.url + '/logout').then(function() {
                $cookies['XSRF-TOKEN'] = undefined;
            });
        };
    }]);
