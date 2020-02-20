
let DatatableRemoteAjax = function () {

    let app = function () {

        let joinText = function(json) {
            let html = '';
            json.order.forEach(k => {
                html += json.detail[k].text + '<br/>';
            })

            return html.substr(0, 100);
        }

        let joinHTML = function(json) {
            let html = '';
            json.order.forEach(k => {
                html += '<i>' + json.detail[k].timeStart + ' - ' + json.detail[k].timeEnd + '</i><br/>' + json.detail[k].text + '</p>';
            })

            return html;
        }

        let getStarHtml = function(num) {
            let html = '';
            for (let i = 0; i < 5; i++) {
                if (num > i) {
                    html += '<i class="fas fa-star"></i>';
                } else {
                    html += '<i class="far fa-star"></i>';
                }
            }

            return html;
        }

        let datatable = $('#data_stats_record').DataTable( {
            "processing": true,
            "serverSide": true,
            "pageLength": 25,
            stateSave: true,
            "drawCallback": function( settings ) {
                $('.modal-record').on('click', function(e){
                    e.preventDefault();

                    let obj = $(this);
                    $.get('/api/record/detail?fid=' + obj.attr('data-id'), function(data){
                        if (data && !data.err) {
                            let obj = $('#modal-default');
                            obj.find(".modal-title").html(data.data.title.substr(0, 50) + '<div style="position: absolute; top: 10px; right: 40px;"><audio controls autoplay><source type="audio/wav" src="' + voiceNoteUrl + '/upload/audios/' + data.data.audioUrl + '">' + '</source></audio></div>');
                            obj.find(".modal-body").html(joinHTML(JSON.parse(data.data.content)));

                            $('#modal-default').modal();
                            $('#modal-default').on('hidden.bs.modal', function (e) {
                                obj.find(".modal-body").html('');
                            })
                        }
                    })
                })
            },
            "ajax": "/api/stats/record" + location.search,
            "columns": [
                { title: 'ID', "data": function(o){
                        return '<a href="' + voiceNoteUrl + '/note/' + o.hashid + '" target="_blank">' + o.id + '</a>'
                }},
                { title: 'User', "data": function(o){
                        return o.user.full_name
                }},
                { title: 'Title', "data": function(o){
                    return '<a href="#" class="modal-record" data-id="' + o.id + '">' + o.title + '</a>'
                }},
                { title: 'Time', width: '40px', "data": function(o){
                        return o.duration
                }},
                { title: 'Content', width: '20%', "data": function(o){
                    try {
                        const json = JSON.parse(o.content);
                        return json.order.length?joinText(json):''
                    } catch (e) {
                        return ''
                    }
                }},
                { title: 'Created', "data": function(o){
                        return moment(o.created).format('YYYY-MM-DD HH:mm A')
                }},
                { title: 'Rating', width: '10%', "data": function(o){
                        return getStarHtml(o.feedback.rate) + (o.feedback.rate?'<p>' + o.feedback.desc + '</p>':'');
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