<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
    <meta name="description" content="Application for following bands.">
    <meta name="author" content="Dmitry Gulyakevich">

    <link rel="icon" type="image/x-icon" href=""/>

    <title>{{.Title}}</title>

    <link href="{{.assetsUrl}}assets/bundle.css" rel="stylesheet">
    <script type="text/javascript" src="{{.assetsUrl}}assets/bundle.js"></script>
</head>
<body>
<header>
    <nav class="navbar navbar-inverse navbar-fixed-top">
        <div class="container">
            <div class="navbar-header">
                <button type="button" class="navbar-toggle collapsed" data-toggle="collapse"
                        data-target="#navbar"
                        aria-expanded="false" aria-controls="navbar">
                    <span class="sr-only">Toggle navigation</span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                </button>

                <a class="navbar-brand" href="/">
                    <i class="fa fa-music fa-1x" aria-hidden="true"></i> Application
                </a>
            </div>
            <div id="navbar" class="navbar-collapse collapse">
                <ul class="nav navbar-nav navbar-right">
                    <li>
                        <p class="navbar-text"></p>
                    </li>
                </ul>
            </div><!--/.navbar-collapse -->
        </div>
    </nav>
</header>

<div class="jumbotron">
    <div class="container">
        Application to test Golang and Vue.Js.
    </div>
</div>

<div class="container">
    <!-- Example row of columns -->
    <div class="row">
        <div class="col-md-9">
            <div id="app"></div>
        </div>
        <div class="col-md-3">
            <div id="sidebar"></div>
        </div>
    </div>

    <hr>
    <footer class="footer">
        <p>&copy; 2017</p>
    </footer>
</div> <!-- /container -->
</body>
</html>
