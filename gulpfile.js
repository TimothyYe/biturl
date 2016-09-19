// Assigning modules to local variables
var gulp = require('gulp');
var sass = require('gulp-sass');
var clean   = require('gulp-clean');
var browserSync = require('browser-sync').create();
var header = require('gulp-header');
var cleanCSS = require('gulp-clean-css');
var rename = require("gulp-rename");
var uglify = require('gulp-uglify');
var pkg = require('./package.json');

// Set the banner content
var banner = ['/*!\n',
    ' * BitURL - <%= pkg.title %> v<%= pkg.version %> (<%= pkg.homepage %>)\n',
    ' * Copyright 2016-' + (new Date()).getFullYear(), ' <%= pkg.author %>\n',
    ' * Licensed under <%= pkg.license.type %> (<%= pkg.license.url %>)\n',
    ' */\n',
    ''
].join('');

gulp.task('sass', function () {
  return gulp.src('./frontend/sass/*.sass')
    .pipe(sass().on('error', sass.logError))
    .pipe(gulp.dest('./app/assets/css'));
});

gulp.task('sass:watch', function () {
  gulp.watch('./frontend/sass/*.scss', ['sass']);
});

// Minify CSS
gulp.task('minify-css', ['clean','sass'], function() {
    return gulp.src('./app/assets/css/style.css')
        .pipe(cleanCSS({ compatibility: 'ie8' }))
        .pipe(header(banner, { pkg: pkg }))
        .pipe(rename({ suffix: '.min' }))
        .pipe(gulp.dest('./app/assets/css'))
        .pipe(browserSync.reload({
            stream: true
        }))
});

// Minify JS
gulp.task('minify-js', ['clean'], function() {
    return gulp.src('./frontend/js/shorten.js')
        .pipe(uglify())
        .pipe(header(banner, { pkg: pkg }))
        .pipe(rename({ suffix: '.min' }))
        .pipe(gulp.dest('./app/assets/js'))
        .pipe(browserSync.reload({
            stream: true
        }))
});

gulp.task('clean', function(){
    return gulp.src([ './app/assets/css/style.min.css', './app/assets/js/shorten.min.js' ], {read: false})
        .pipe(clean());
});

gulp.task('clean-unused-file', ['minify-css'], function(){
    return gulp.src(['./app/assets/css/style.css'], {read: false})
        .pipe(clean());
});

gulp.task('build', ['clean','sass','minify-css','minify-js','clean-unused-file']);
