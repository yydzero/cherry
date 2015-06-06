'use strict';

/**
 * @ngdoc function
 * @name clientApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the clientApp
 */
angular.module('clientApp').controller('MainCtrl', ['$scope', '$http', '$log', 'config', 'gzhService', function ($scope, $http, $log, config, gzhService) {

    $scope.message = '';

    $scope.newSubscription = '';

    $scope.subscripts = [];

    $scope.config = config;

    console.log($scope.config);

    $scope.getSubscriptions = function() {
        $http.get(config.url + '/gzh')
            .success(function (resp) {
                $log.info(resp);
                if (resp.Status === 'ok') {
                    $scope.subscripts = resp.Resource;
                    $scope.message = '';
                } else {
                    $scope.message = resp.Message;
                }
            }).error(function (error) {
                $scope.message = '不能获得订阅的公众号列表.';
            }
        );
    };

    $scope.addSubscription = function() {
        console.log('subscript is called.');
        var ns = {
            name: $scope.newSubscription.trim()
        };
        if (! ns.name) {
            return;
        }

        // 验证公众号是否存在
        $scope.message = '查找公众号：' + $scope.newSubscription + '  ...';

        $http.post(config.url + '/gzh', { 'Name': ns.name })
            .success(function (resp) {
                $log.info(resp);
                if (resp.Status === 'ok') {
                    $scope.subscripts = $scope.subscripts.concat(resp.Resource);
                    $scope.newSubscription = '';
                } else {
                    $scope.message = resp.Message;
                }
            }).error(function (error) {
                $scope.message = '公众号：' + $scope.newSubscription + ' 不存在';
            }
        );
    };

    $scope.removeSubscription = function(subscription) {
        console.log(subscription);
        if (!subscription || !subscription.OpenId) {
            return;
        }

        console.log('remove is called');

        $http.delete(config.url + '/gzh/' + subscription.OpenId)
            .success(function (resp) {
                $log.info(resp);
                if (resp.Status === 'ok') {
                    _.remove($scope.subscripts, function (v) { return v === subscription; });
                    $scope.newSubscription = '';
                } else {
                    $scope.message = resp.Message;
                }
            }).error(function (error) {
                $scope.message = '公众号：' + $scope.newSubscription + ' 不存在';
            }
        );
    };

    $scope.getSubscriptions();

}]);
