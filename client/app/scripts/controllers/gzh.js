'use strict';


angular.module('clientApp').controller('GzhCtrl', ['$scope', '$http', '$routeParams', '$log', 'config', function ($scope, $http, $routeParams, $log, config) {

    $scope.articles = [];
    $scope.message = '';
    $scope.deleting = false;

    $scope.getArticlesByOpenId = function(openId) {
        console.log(openId);
        $http.get(config.url + '/article?openId=' + openId)
            .success(function (resp) {
                $log.info("article with openId: ", resp);
                if (resp.Status === 'ok') {
                    $scope.articles = resp.Resource;
                    $scope.message = '';
                } else {
                    $scope.message = resp.Message;
                }
            }).error(function (error) {
                $scope.message = '不能获得订阅的公众号列表.';
            }
        );
    };

    $scope.removeArticle = function(article) {
        $scope.deleting = true;
        console.log(article);
        $http.delete(config.url + '/article/' + article.Id)
            .success(function (resp) {
                $log.info("delete article response", resp);
                if (resp.Status === 'ok') {
                    _.remove($scope.articles, function (v) { return v.Id === article.Id; });
                    $scope.message = '';
                } else {
                    $scope.message = resp.Message;
                    console.log("delete article " + article.Id + " :" + $scope.message);
                }
                $scope.deleting = false;
            }).error(function (error) {
                $scope.message = '不能获得订阅的公众号列表.';
                console.log("delete article " + article.Id + " :" + $scope.message);
                $scope.deleting = false;
            }
        );
    };

    $scope.getArticlesByOpenId($routeParams.openId);
}]);
