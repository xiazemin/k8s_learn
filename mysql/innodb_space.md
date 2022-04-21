% innodb_space -s /opt/homebrew/var/mysql/ibdata1 system-spaces
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777
name                            pages       indexes
(system)                        768         7
mysql/engine_cost               6           1
mysql/gtid_executed             6           1
mysql/help_category             7           2
mysql/help_keyword              18          2
mysql/help_relation             10          1
mysql/help_topic                576         2
mysql/innodb_index_stats        6           1
mysql/innodb_table_stats        6           1
mysql/plugin                    6           1
mysql/server_cost               6           1
mysql/servers                   6           1
mysql/slave_master_info         6           1
mysql/slave_relay_log_info      6           1
mysql/slave_worker_info         6           1
mysql/time_zone                 6           1
mysql/time_zone_leap_second     6           1
mysql/time_zone_name            6           1
mysql/time_zone_transition      6           1
mysql/time_zone_transition_type 6           1
svc_tree/edge                   7           2
sys/sys_config                  6           1


https://github.com/jeremycole/innodb_ruby/wiki
https://github.com/jeremycole/innodb_ruby
系统空间里创建了7个页面


% innodb_space -s /opt/homebrew/var/mysql/ibdata1  -T svc_tree/edge space-indexes
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777
id          name                            root        fseg        fseg_id     used        allocated   fill_factor 
41          PRIMARY                         3           internal    1           1           1           100.00%     
41          PRIMARY                         3           leaf        2           0           0           0.00%       
42          node_parent_link                4           internal    3           1           1           100.00%     
42          node_parent_link                4           leaf        4           0           0           0.00%       

主索引和聚簇索引分别分配了一个内部区，没有叶子索引没有叶子分区

 % innodb_space -s /opt/homebrew/var/mysql/ibdata1  -T svc_tree/edge space-page-type-regions
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777
start       end         count       type                
0           0           1           FSP_HDR             
1           1           1           IBUF_BITMAP         
2           2           1           INODE               
3           4           2           INDEX               
5           6           2           FREE (ALLOCATED)   


% innodb_space -s /opt/homebrew/var/mysql/ibdata1  -T svc_tree/edge space-page-type-summary
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777
type                count       percent     description         
INDEX               2           28.57       B+Tree index        
ALLOCATED           2           28.57       Freshly allocated   
FSP_HDR             1           14.29       File space header   
IBUF_BITMAP         1           14.29       Insert buffer bitmap
INODE               1           14.29       File segment inode  


% innodb_space -s /opt/homebrew/var/mysql/ibdata1  -T svc_tree/edge space-extents-illustrate
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777

                                     svc_tree/edge.ibd                         
  Start Page ╭────────────────────────────────────────────────────────────────╮
           0 │███▁▁░░                                                         │
             ╰────────────────────────────────────────────────────────────────╯

Legend (█ = 1 page):
  Page Type                                                         Pages    Ratio
  █ System                                                              3   42.86%
  █ Index 41 (svc_tree/edge.PRIMARY)                                    1   14.29%
  █ Index 42 (svc_tree/edge.node_parent_link)                           1   14.29%
  ░ Free space                                                          2   28.57%

% innodb_space -s /opt/homebrew/var/mysql/ibdata1  -T svc_tree/edge space-lsn-age-illustrate
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777

                                     svc_tree/edge.ibd                         
  Start Page ╭────────────────────────────────────────────────────────────────╮
           0 │█████  │
             ╰────────────────────────────────────────────────────────────────╯

LSN Age Histogram (█ = ~0 pages):
       Min LSN ▃           ▃                             █ Max LSN     
       2754749 ███████████████████████████████████████████ 2756293   


% innodb_space -s /opt/homebrew/var/mysql/ibdata1  -T svc_tree/edge -p 3 page-account       
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777
Accounting for page 3:
  Page type is INDEX (B+Tree index, table and index data stored in B+Tree structure).
  Extent descriptor for pages 0-63 is at page 0, offset 158.
  Extent is not fully allocated to an fseg; may be a fragment extent.
  Page is marked as used in extent descriptor.
  Extent is in free_frag list of space.
  Page is in fragment array of fseg 1.
  Fseg is in internal fseg of index 41.
  Index root is page 3.
  Index is svc_tree/edge.PRIMARY.


 % innodb_space -s /opt/homebrew/var/mysql/ibdata1  -T svc_tree/edge -p 3 page-dump    
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777
#<Innodb::Page::Index:0x00007fe6d08c9ab0>:

fil header:
#<struct Innodb::Page::FilHeader
 checksum=818688513,
 offset=3,
 prev=nil,
 next=nil,
 lsn=2754749,
 type=:INDEX,
 flush_lsn=0,
 space_id=23>

fil trailer:
#<struct Innodb::Page::FilTrailer checksum=818688513, lsn_low32=2754749>

page header:
#<struct Innodb::Page::Index::PageHeader
 n_dir_slots=2,
 heap_top=120,
 n_heap_format=32770,
 n_heap=2,
 format=:compact,
 garbage_offset=0,
 garbage_size=0,
 last_insert_offset=0,
 direction=:no_direction,
 n_direction=0,
 n_recs=0,
 max_trx_id=0,
 level=0,
 index_id=41>

fseg header:
#<struct Innodb::Page::Index::FsegHeader
 leaf=
  <Innodb::Inode space=<Innodb::Space file="svc_tree/edge.ibd", page_size=16384, pages=7>, fseg=2>,
 internal=
  <Innodb::Inode space=<Innodb::Space file="svc_tree/edge.ibd", page_size=16384, pages=7>, fseg=1>>

sizes:
  header           120
  trailer            8
  directory          4
  free           16252
  used             132
  record             0
  per record      0.00

page directory:
[99, 112]

system records:
#<struct Innodb::Page::Index::SystemRecord
 offset=99,
 header=
  #<struct Innodb::Page::Index::RecordHeader
   length=5,
   next=112,
   type=:infimum,
   heap_number=0,
   n_owned=1,
   info_flags=0,
   offset_size=nil,
   n_fields=nil,
   nulls=nil,
   lengths=nil,
   externs=nil>,
 next=112,
 data="infimum\x00",
 length=8>
#<struct Innodb::Page::Index::SystemRecord
 offset=112,
 header=
  #<struct Innodb::Page::Index::RecordHeader
   length=5,
   next=112,
   type=:supremum,
   heap_number=1,
   n_owned=1,
   info_flags=0,
   offset_size=nil,
   n_fields=nil,
   nulls=nil,
   lengths=nil,
   externs=nil>,
 next=112,
 data="supremum",
 length=8>

garbage records:

records:


 % innodb_space -s /opt/homebrew/var/mysql/ibdata1  -T svc_tree/edge -p 3 page-records
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777


% innodb_space -s /opt/homebrew/var/mysql/ibdata1  -T svc_tree/edge -p 3 page-directory-summary
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777
slot    offset  type          owned   key
0       99      infimum       1       
1       112     supremum      1       



% innodb_space -s /opt/homebrew/var/mysql/ibdata1  -T svc_tree/edge -p 3 page-illustrate       
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777

                                       Page 3 (INDEX)                          
      Offset ╭────────────────────────────────────────────────────────────────╮
           0 │█████████████████████████████████████▋██████████████████████████│
          64 │█████████▋███████████████████▋████████████▋████████████▋        │
         ... │                                                                │
       16320 │                                                      █▋█▋█████▋│
             ╰────────────────────────────────────────────────────────────────╯

Legend (█ = 1 byte):
  Region Type                         Bytes    Ratio
  █ FIL Header                           38    0.23%
  █ Index Header                         36    0.22%
  █ File Segment Header                  20    0.12%
  █ Infimum                              13    0.08%
  █ Supremum                             13    0.08%
  █ Page Directory                        4    0.02%
  █ FIL Trailer                           8    0.05%
  ░ Garbage                               0    0.00%
    Free                              16252   99.19%


% innodb_space -s /opt/homebrew/var/mysql/ibdata1  -T svc_tree/edge -I PRIMARY index-recurse
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777
ROOT NODE #3: 0 records, 0 bytes

 % innodb_space -s /opt/homebrew/var/mysql/ibdata1  -T svc_tree/edge -I PRIMARY index-record-offsets
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777
page_offset         record_offset       


% innodb_space -s /opt/homebrew/var/mysql/ibdata1  -T svc_tree/edge -I PRIMARY -l 0 index-level-summary
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777
page    index   level   data    free    records min_key 
/Library/Ruby/Gems/2.6.0/gems/innodb_ruby-0.12.0/bin/innodb_space:1210:in `block in index_level_summary': undefined method `key_string' for nil:NilClass (NoMethodError)
        from /Library/Ruby/Gems/2.6.0/gems/innodb_ruby-0.12.0/lib/innodb/index.rb:147:in `each_page_from'
        from /Library/Ruby/Gems/2.6.0/gems/innodb_ruby-0.12.0/lib/innodb/index.rb:159:in `each_page_at_level'
        from /Library/Ruby/Gems/2.6.0/gems/innodb_ruby-0.12.0/bin/innodb_space:1202:in `index_level_summary'
        from /Library/Ruby/Gems/2.6.0/gems/innodb_ruby-0.12.0/bin/innodb_space:1760:in `<top (required)>'
        from /usr/local/bin/innodb_space:23:in `load'
        from /usr/local/bin/innodb_space:23:in `<main>'


 % innodb_space -s /opt/homebrew/var/mysql/ibdata1  -T svc_tree/edge -p 3 -R 128 record-dump            
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777
Record at offset 128

Header:
  Next record offset  : 128
  Heap number         : 0
  Type                : conventional
  Deleted             : false
  Length              : 5

System fields:
  Transaction ID: 0
  Roll Pointer:
    Undo Log: page 0, offset 0
    Rollback Segment ID: 0
    Insert: false

Key fields:
  parent_id: -9223372036854775808
  parent_type: -128
  node_id: -9223372036854775808
  node_type: -128

Non-key fields:
  version: -9223372036854775808
  order: -9223372036854775808
  is_link: "0b0"
  is_removed: "0b0"

% innodb_space -s /opt/homebrew/var/mysql/ibdata1  -T svc_tree/edge -p 3 -R 128 record-history
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777
Transaction   Type                Undo record
/Library/Ruby/Gems/2.6.0/gems/bindata-2.4.10/lib/bindata/io.rb:316:in `read': End of file reached (EOFError)
        from /Library/Ruby/Gems/2.6.0/gems/bindata-2.4.10/lib/bindata/io.rb:278:in `readbytes'


% innodb_space -s /opt/homebrew/var/mysql/ibdata1  space-lists                                
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777
name                length      f_page      f_offset    l_page      l_offset    
free                0           0           0           0           0           
free_frag           1           0           318         0           318         
full_frag           2           0           158         0           278         
full_inodes         1           2           38          2           38          
free_inodes         1           243         38          243         38    


% innodb_space -s /opt/homebrew/var/mysql/ibdata1  space-list-iterate -L free_frag
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777
start_page  page_used_bitmap                                                
256         ##########################################################......



% innodb_space -s /opt/homebrew/var/mysql/ibdata1  space-inodes-summary           
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777
INODE fseg_id=1, pages=2, frag=2, full=0, not_full=0, free=0
INODE fseg_id=2, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=3, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=4, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=5, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=6, pages=0, frag=0, full=0, not_full=0, free=0
INODE fseg_id=7, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=8, pages=0, frag=0, full=0, not_full=0, free=0
INODE fseg_id=9, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=10, pages=0, frag=0, full=0, not_full=0, free=0
INODE fseg_id=11, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=12, pages=0, frag=0, full=0, not_full=0, free=0
INODE fseg_id=13, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=14, pages=0, frag=0, full=0, not_full=0, free=0
INODE fseg_id=15, pages=160, frag=32, full=2, not_full=0, free=0
INODE fseg_id=16, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=17, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=18, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=19, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=20, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=21, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=22, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=23, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=24, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=25, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=26, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=27, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=28, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=29, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=30, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=31, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=32, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=33, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=34, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=35, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=36, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=37, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=38, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=39, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=40, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=41, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=42, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=43, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=44, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=45, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=46, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=47, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=48, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=49, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=50, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=51, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=52, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=53, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=54, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=55, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=56, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=57, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=58, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=59, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=60, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=61, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=62, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=63, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=64, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=65, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=66, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=67, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=68, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=69, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=70, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=71, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=72, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=73, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=74, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=75, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=76, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=77, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=78, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=79, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=80, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=81, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=82, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=83, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=84, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=85, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=86, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=87, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=88, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=89, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=90, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=91, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=92, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=93, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=94, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=95, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=96, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=97, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=98, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=99, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=100, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=101, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=102, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=103, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=104, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=105, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=106, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=107, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=108, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=109, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=110, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=111, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=112, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=113, pages=0, frag=0, full=0, not_full=0, free=0
INODE fseg_id=114, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=115, pages=0, frag=0, full=0, not_full=0, free=0
INODE fseg_id=116, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=117, pages=0, frag=0, full=0, not_full=0, free=0
INODE fseg_id=118, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=119, pages=0, frag=0, full=0, not_full=0, free=0
INODE fseg_id=120, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=121, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=122, pages=0, frag=0, full=0, not_full=0, free=0
INODE fseg_id=123, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=124, pages=0, frag=0, full=0, not_full=0, free=0
INODE fseg_id=125, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=126, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=127, pages=0, frag=0, full=0, not_full=0, free=0
INODE fseg_id=128, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=129, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=130, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=131, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=132, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=133, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=134, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=135, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=136, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=137, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=138, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=139, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=140, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=141, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=142, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=143, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=144, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=145, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=146, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=147, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=148, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=149, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=150, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=151, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=152, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=153, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=154, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=155, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=156, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=157, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=158, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=160, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=161, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=162, pages=1, frag=1, full=0, not_full=0, free=0
INODE fseg_id=163, pages=1, frag=1, full=0, not_full=0, free=0


% innodb_space -s /opt/homebrew/var/mysql/ibdata1  undo-history-summary
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777
Page    Offset  Transaction   Type                Table

