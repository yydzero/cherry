'use strict';

/**
 * @ngdoc function
 * @name clientApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the clientApp
 */
angular.module('clientApp').controller('CrawlerCtrl', ['$scope', 'gzhService', function ($scope, gzhService) {

    $scope.isLoading = gzhService.isCrawlLoading;
    $scope.crawlMessage = gzhService.crawlMessage;
    $scope.crawlStats = gzhService.crawlStats;

    gzhService.getGzhCrawlStatus().then(function(data) {
        console.log("in then:");
        $scope.crawlStats = data;
    });

    $scope.crawlGzh = function(gzh) {
        console.log(gzh);

        gzhService.crawlGzh(gzh).then(function(data) {
            console.log(data);
        })
    };

    //$scope.isLoading = false;
    //$scope.crawlMessage = '';
    //$scope.crawlStats = [];
    //
    //var getStats = function () {
    //    gzhService.getGzhCrawlStatus().success(function (resp) {
    //        $log.info(resp);
    //        if (resp.Status === 'ok') {
    //            $scope.crawlStats = resp.Resource;
    //            $scope.crawlMessage = '';
    //        } else {
    //            $scope.crawlMessage = resp.Message
    //        }
    //        $scope.isLoading = false;
    //    }).error(function (error) {
    //        $scope.crawlMessage = '不能获得订阅的公众号爬虫状态信息.';
    //        $scope.isLoading = false;
    //    });
    //};
    //
    //getStats();
}]);
