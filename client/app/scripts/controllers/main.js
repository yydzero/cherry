'use strict';

/**
 * @ngdoc function
 * @name clientApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the clientApp
 */
angular.module('clientApp').controller('MainCtrl', ['$scope', '$http', '$log', 'config', function ($scope, $http, $log, config) {

    $scope.subscripts = [
        {
            name: '大数据文摘'
        },
        {
            name: '大数据那些事'
        }
    ];

    $scope.newSubscription = '';

    $scope.addSubscription = function() {
        console.log('subscript is called.');
        var ns = {
            name: $scope.newSubscription.trim()
        };
        if (! ns.name) {
            return
        }

        // 验证公众号是否存在
        $scope.newSubscription = '验证公众号：' + $scope.newSubscription + ' 是否存在 ...';



        $scope.subscripts.push(ns)
    };


}]);
