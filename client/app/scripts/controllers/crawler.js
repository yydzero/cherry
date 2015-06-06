'use strict';

/**
 * @ngdoc function
 * @name clientApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the clientApp
 */
angular.module('clientApp').controller('CrawlerCtrl', ['$scope', '$timeout', 'gzhService', function ($scope, $timeout, gzhService) {

    $scope.isLoading = gzhService.isCrawlLoading;
    $scope.crawlMessage = gzhService.crawlMessage;
    $scope.crawlStats = gzhService.crawlStats;

    var timeout;

    gzhService.getGzhCrawlStatus().then(function(data) {
        console.log("in then:");
        $scope.crawlStats = data;
    });

    $scope.crawlGzh = function(gzh) {
        $scope.crawlMessage = '抓取公众号：' + gzh.Title + ' ...';
        console.log(gzh);

        gzhService.crawlGzh(gzh).then(function(data) {
            console.log(data);
            $scope.crawlMessage = '抓取公众号：' + gzh.Title + ' 完成';

            if (timeout) $timeout.cancel(timeout);
            timeout = $timeout(function() {
                $scope.crawlMessage = '';
                $scope.$apply($scope.crawlMessage);
            }, 5000);
        });
    };

    // crawl all gzh
    $scope.refreshAll = function() {
        $scope.crawlMessage = '刷新所有公众号中, 请耐心等待一段时间再刷新页面，耗时较长 ...';

        gzhService.crawlAll().then(function(data) {
            console.log(data);

            $scope.crawlMessage = '刷新所有公众号：' + gzh.Title + ' 完成';

            if (timeout) $timeout.cancel(timeout);
            timeout = $timeout(function() {
                $scope.crawlMessage = '';
                $scope.$apply($scope.crawlMessage);
            }, 5000);
        });
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
