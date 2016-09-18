// Assigning modules to local variables
var gulp = require('gulp');
var sass = require('gulp-sass');
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
  return gulp.src('./frontend/sass/*.scss')
    .pipe(sass().on('error', sass.logError))
    .pipe(gulp.dest('./app/assets/css'));
});

gulp.task('sass:watch', function () {
  gulp.watch('./frontend/sass/*.scss', ['sass']);
});

// Minify CSS
gulp.task('minify-css', ['sass'], function() {
    return gulp.src('./app/assets/css/style.css')
        .pipe(cleanCSS({ compatibility: 'ie8' }))
        .pipe(rename({ suffix: '.min' }))
        .pipe(gulp.dest('./app/assets/css'))
        .pipe(browserSync.reload({
            stream: true
        }))
});

gulp.task('build', ['sass','minify-css']);
