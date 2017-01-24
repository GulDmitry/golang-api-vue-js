<div class="row">
    <div id="main" class="col-sm-9">

        <h1> Edit Task </h1>

        <form name="post" method="post">
            <div class="form-group">
                <label class="control-label required" for="post_title">Title</label>
                <input value="{{.task.Title}}" type="text" id="post_title" name="title" required="required" autofocus="autofocus" class="form-control">
            </div>
            <div class="form-group">
                <label class="control-label required" for="post_body">Body</label>
                <textarea id="post_body" name="body" required="required" rows="10" class="form-control">{{.task.Body}}</textarea>
            </div>

            <input type="submit" value="Save Task" class="btn btn-primary">

            <a href="/" class="btn btn-link">
                Back to list
            </a>
        </form>
    </div>
</div>