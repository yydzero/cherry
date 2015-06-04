'use strict';

/**
 * @ngdoc function
 * @name clientApp.controller:MainCtrl
 * @description
 * # MainCtrl
 * Controller of the clientApp
 */
angular.module('clientApp').controller('MainCtrl', ['$scope', function ($scope) {

    // below are deprecated

    $scope.info = '';
    $scope.groups = [];

    $scope.getGroups = function () {
        $http.get(config.url + '/groups')
            .success(function (resp) {
                $scope.groups = resp.Resource;
                $log.info($scope.groups);
            }).error(function (error) {
                $log.info('group save failed');
            });
    };

    $scope.getGroups();

    $scope.addGroup = function () {
        if ($scope.groups.length > 10) {
            window.alert('You can\'t add more than 10 groups!');
            return;
        }
        var groupName = document.getElementById('groupName').value;
        if (groupName.length > 0) {
            $http.post(config.url + '/groups', {
                name: groupName,
                categories: ''
            }).success(function () {
                $log.info('group saved correctly');

                $scope.getGroups();
            }).error(function (error) {
                $log.info('group save failed');
            });

            document.getElementById('groupName').value = '';
        }
    };

    $scope.editGroup = function (group) {
        group.editing = true;
    };

    $scope.cancelEditingGroup = function (group) {
        group.editing = false;
    };

    $scope.saveGroup = function (group) {
        group.save();
    };

    $scope.removeGroup = function (group) {
        if (window.confirm('Are you sure to remove this group?')) {
            var index = $scope.groups.indexOf(group);
            $scope.groups.splice(index, 1);

            $http.delete(config.url + '/groups/' + group.Id).success(function () {
                $log.info('group delete correctly');
            }).error(function (error) {
                $log.info('group delete failed');
            });
        }
    };

    $scope.saveGroups = function () {
        for (var i = $scope.groups.length - 1; i >= 0; i--) {
            var group = $scope.groups[i];
            group.sortOrder = i + 1;
            group.save();
        }
    };

    $scope.addCategory = function (group) {
        if (!group.newCategoryName || group.newCategoryName.length === 0) {
            return;
        }
        group.categories.$add({
            name: group.newCategoryName,
            sortOrder: group.categories.length,
            type: 'category'
        });
        group.newCategoryName = '';
        group.save();
    };

    $scope.removeCategory = function (group, category) {
        if (window.confirm('Are you sure to remove this category?')) {
            var index = group.categories.indexOf(category);
            if (index > -1) {
                group.categories.splice(index, 1)[0];
            }
            group.save();
        }
    };

    $scope.options = {
        accept: function (sourceNode, destNodes, destIndex) {
            var data = sourceNode.$modelValue;
            var destType = destNodes.$element.attr('data-type');
            return (data.type == destType); // only accept the same type
        },
        dropped: function (event) {
            console.log(event);
            var sourceNode = event.source.nodeScope;
            var destNodes = event.dest.nodesScope;
            // update changes to server
            if (destNodes.isParent(sourceNode)
                && destNodes.$element.attr('data-type') == 'category') { // If it moves in the same group, then only update group
                var group = destNodes.$nodeScope.$modelValue;
                group.save();
            } else { // save all
                $scope.saveGroups();
            }
        }
    };

    $scope.list = [
        {
            'title': 'group1',
            'items': [
                {'title': 'dezhou'},
                {'title': 'taian'}
            ]
        }
    ];

    // ------------------

    $scope.remove = function(scope) {
        scope.remove();
    };

    $scope.toggle = function(scope) {
        console.log('toggle this');
        scope.toggle();
    };

    $scope.moveLastToTheBeginning = function () {
        var a = $scope.data.pop();
        $scope.data.splice(0,0, a);
    };

    $scope.newSubItem = function(scope) {
        var nodeData = scope.$modelValue;
        nodeData.nodes.push({
            id: nodeData.id * 10 + nodeData.nodes.length,
            title: nodeData.title + '.' + (nodeData.nodes.length + 1),
            nodes: []
        });
    };

    var getRootNodesScope = function() {
        return angular.element(document.getElementById('tree-root')).scope();
    };

    $scope.collapseAll = function() {
        var scope = getRootNodesScope();
        scope.collapseAll();
    };

    $scope.expandAll = function() {
        var scope = getRootNodesScope();
        scope.expandAll();
    };

    $scope.data = [{
        'id': 1,
        'title': 'node1',
        'nodes': [
            {
                'id': 11,
                'title': 'node1.1',
                'nodes': [
                    {
                        'id': 111,
                        'title': 'node1.1.1',
                        'nodes': []
                    }
                ]
            },
            {
                'id': 12,
                'title': 'node1.2',
                'nodes': []
            }
        ],
    }, {
        'id': 2,
        'title': 'node2',
        'nodes': [
            {
                'id': 21,
                'title': 'node2.1',
                'nodes': []
            },
            {
                'id': 22,
                'title': 'node2.2',
                'nodes': []
            }
        ],
    }, {
        'id': 3,
        'title': 'node3',
        'nodes': [
            {
                'id': 31,
                'title': 'node3.1',
                'nodes': []
            }
        ]
    }];

    $scope.myData = [
        {
            'firstName': 'Cox',
            'lastName': 'Carney',
            'company': 'Enormo',
            'employed': true
        },
        {
            'firstName': 'Lorraine',
            'lastName': 'Wise',
            'company': 'Comveyer',
            'employed': false
        },
        {
            'firstName': 'Nancy',
            'lastName': 'Waters',
            'company': 'Fuelton',
            'employed': false
        }
    ];

    $scope.list = [
        {
            'title': 'group1',
            'items': [
                { 'title': 'dezhou' },
                { 'title': 'taian' }
            ]
        }
    ]
}]);
