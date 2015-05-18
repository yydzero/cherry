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
                    if (data.Status === 'ok') {
                        console.log(data);
                        //$cookies['XSRF-TOKEN'] = data.Resource.sessionid;
                        $cookies['USER'] = data.Resource;

                        console.log('signin is finished with ' + data);

                        //$state.go('projects');
                        $location.path('/');
                    } else {
                        $scope.message = 'failed: ' + data.Message;
                    }



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
