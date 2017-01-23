<div class="row">
    <div id="main" class="col-sm-9">

        <h1> New Task </h1>

        <form name="post" method="post">
            <div class="form-group">
                <label class="control-label required" for="post_title">Title</label>
                <input type="text" id="post_title" name="title" required="required" autofocus="autofocus" class="form-control">
            </div>
            <div class="form-group">
                <label class="control-label required" for="post_body">Body</label>
                <textarea id="post_body" name="body" required="required" rows="20" class="form-control"></textarea>
            </div>
            {{/*<div class="form-group">
                <label class="control-label required" for="post_publishedAt">Published at</label>
                <div class="input-group date" data-toggle="datetimepicker">
                    <input type="datetime" id="post_publishedAt" name="post[publishedAt]" required="required" data-date-format="YYYY-MM-DDTHH:mm:ssZ" data-date-locale="en" class="form-control" value="2017-01-23T13:16:48+03:00">
                    <span class="input-group-addon"><span class="fa fa-calendar" aria-hidden="true"></span></span>
                </div>
            </div>*/}}

            <input type="submit" value="Create Task" class="btn btn-primary">

            <a href="/" class="btn btn-link">
                Back to list
            </a>
        </form>
    </div>
</div>