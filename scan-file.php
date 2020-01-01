<?php
/**
 * 遍历文件夹下的所有文件
 * @param $dir
 */
function scan($dir) {
    $files = scandir($dir);
    foreach ($files as $file) {
        $next = sprintf('%s/%s', $dir, $file);
        // 如果是文件夹，递归扫描
        if (is_dir($next)) {
            if ($file == '.' || $file == '..') {
                continue;
            }
            // echo $next . PHP_EOL;
            scan($next);
        } else {
            echo $next . PHP_EOL;
        }
    }
}
