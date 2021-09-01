 https://github.com/wp-statistics/GeoLite2-City

 curl -O https://raw.githubusercontent.com/wp-statistics/GeoLite2-City/master/GeoLite2-City.mmdb.gz

% tar -zxvf GeoLite2-City.mmdb.gz 
tar: Error opening archive: Unrecognized archive format

% gzip -d  GeoLite2-City.mmdb.gz 
gzip: GeoLite2-City.mmdb.gz: unexpected end of file
gzip: GeoLite2-City.mmdb.gz: uncompress failed

go mod init geoip

https://github.com/oschwald/maxminddb-golang