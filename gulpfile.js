var gulp = require('gulp');
var less = require('gulp-less');
var changed = require("gulp-changed");
var imagemin = require('gulp-imagemin');
var uglifycss = require("gulp-uglifycss");
var minimizejs = require("gulp-uglify");
var sequence = require("run-sequence");
var watch = require("gulp-watch");
var typescript = require("gulp-typescript");

var release = false;

gulp.task('assets:favicons', function() {
	gulp.src("assets/favicons/*.ico")
		.pipe(changed("static"))
		.pipe(gulp.dest("static"));

	gulp.src("assets/favicons/*.png")
		.pipe(changed("static"))
		.pipe(imagemin({
			progressive: false	
		}))
		.pipe(gulp.dest("static"));
})

gulp.task('assets:less', function() {
	gulp.src("assets/less/*.less")
		.pipe(changed("static/css", { extension: ".css" }))
		.pipe(less())
		.pipe(gulp.dest("static/css"));
})

gulp.task('assets:typescript', function() {
	var stream = gulp.src("assets/ts/*.ts")
		.pipe(typescript())
		.pipe(gulp.dest("static/js"));

	stream.on("error", function() {
		console.log(arguments);
	});
})

gulp.task('flag:release', function() {
	release = true;
})

gulp.task('assets:minimizecss', function() {
	gulp.src("static/css/*.css")
		.pipe(uglifycss())
		.pipe(gulp.dest("static/css"));
})

gulp.task('assets:minimizejs', function() {
	gulp.src("static/js/*.js")
		.pipe(minimizejs())
		.pipe(gulp.dest("static/js"));
})

gulp.task('default', ['assets:favicons', 'assets:less', 'assets:typescript']);
gulp.task('release', function() {
	sequence('flag:release', ['default'], ['assets:minimizecss', 'assets:minimizejs']);
});

gulp.task('watch', function() {
	watch("assets/less/*.less", function() {
		gulp.start("assets:less");
	});

	watch("assets/ts/*.ts", function() {
		gulp.start("assets:typescript");
	});
})