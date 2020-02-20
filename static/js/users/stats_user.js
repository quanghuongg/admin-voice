
let DatatableRemoteAjax = function () {

    let app = function () {

        let datatable = $('#data_stats_user').DataTable( {
            "processing": true,
            "serverSide": true,
            "pageLength": 25,
            stateSave: true,
            "ajax": "/api/stats/user",
            "columns": [
                { title: 'ID', "data": function(o){
                    return o.id
                }},
                { title: 'Avatar', "data": function(o){
                        return '<img style="height: 50px;" src="' + o.avatar + '" alt="" />'
                    }},
                { title: 'Fullname', "data": function(o){
                    return o.full_name
                }},
                { title: 'Email', "data": function(o){
                    return o.email
                }},
                { title: 'Phone', "data": function(o){
                    return o.phone
                }},
                { title: 'Created', "data": function(o){
                    return moment(o.created).format('YYYY-MM-DD HH:mm A')
                }},
                { title: 'Records', width: '40px', "data": function(o){
                    return '<a target="_blank" href="/stats/record?uid=' + o.id + '">' + o.totalRecords + '</a>'
                }},
            ],
        } );

    };

    return {
        init: function () {
            app();
        }
    };
}();

$(function () {
    DatatableRemoteAjax.init();
});