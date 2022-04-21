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

 % innodb_space -s /opt/homebrew/var/mysql/ibdata1  -T svc_tree/edge space-indexes
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777
id          name                            root        fseg        fseg_id     used        allocated   fill_factor
41          PRIMARY                         3           internal    1           1           1           100.00%
41          PRIMARY                         3           leaf        2           0           0           0.00%
42          node_parent_link                4           internal    3           1           1           100.00%
42          node_parent_link                4           leaf        4           0           0           0.00%



% innodb_space -s /opt/homebrew/var/mysql/ibdata1  -T svc_tree/edge -p 3 page-dump 
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777
#<Innodb::Page::Index:0x00007fe4600c9328>:

fil header:
#<struct Innodb::Page::FilHeader
 checksum=2649160932,
 offset=3,
 prev=nil,
 next=nil,
 lsn=2759373,
 type=:INDEX,
 flush_lsn=0,
 space_id=23>

fil trailer:
#<struct Innodb::Page::FilTrailer checksum=2649160932, lsn_low32=2759373>

page header:
#<struct Innodb::Page::Index::PageHeader
 n_dir_slots=2,
 heap_top=174,
 n_heap_format=32771,
 n_heap=3,
 format=:compact,
 garbage_offset=0,
 garbage_size=0,
 last_insert_offset=125,
 direction=:no_direction,
 n_direction=0,
 n_recs=1,
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
  free           16198
  used             186
  record            54
  per record     54.00

page directory:
[99, 112]

system records:
#<struct Innodb::Page::Index::SystemRecord
 offset=99,
 header=
  #<struct Innodb::Page::Index::RecordHeader
   length=5,
   next=125,
   type=:infimum,
   heap_number=0,
   n_owned=1,
   info_flags=0,
   offset_size=nil,
   n_fields=nil,
   nulls=nil,
   lengths=nil,
   externs=nil>,
 next=125,
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
   n_owned=2,
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
#<struct Innodb::Page::Index::UserRecord
 type=:clustered,
 format=:compact,
 offset=125,
 header=
  #<struct Innodb::Page::Index::RecordHeader
   length=5,
   next=112,
   type=:conventional,
   heap_number=2,
   n_owned=0,
   info_flags=0,
   offset_size=nil,
   n_fields=nil,
   nulls=[],
   lengths={},
   externs=[]>,
 next=112,
 key=
  [#<struct Innodb::Page::Index::FieldDescriptor
    name="parent_id",
    type="BIGINT",
    value=0,
    extern=nil>,
   #<struct Innodb::Page::Index::FieldDescriptor
    name="parent_type",
    type="TINYINT",
    value=8,
    extern=nil>,
   #<struct Innodb::Page::Index::FieldDescriptor
    name="node_id",
    type="BIGINT",
    value=900000000005414,
    extern=nil>,
   #<struct Innodb::Page::Index::FieldDescriptor
    name="node_type",
    type="TINYINT",
    value=14,
    extern=nil>],
 row=
  [#<struct Innodb::Page::Index::FieldDescriptor
    name="version",
    type="BIGINT",
    value=1553767140227,
    extern=nil>,
   #<struct Innodb::Page::Index::FieldDescriptor
    name="order",
    type="BIGINT",
    value=1551088328795,
    extern=nil>,
   #<struct Innodb::Page::Index::FieldDescriptor
    name="is_link",
    type="BIT UNSIGNED",
    value="0b1",
    extern=nil>,
   #<struct Innodb::Page::Index::FieldDescriptor
    name="is_removed",
    type="BIT UNSIGNED",
    value="0b0",
    extern=nil>],
 sys=
  [#<struct Innodb::Page::Index::FieldDescriptor
    name="DB_TRX_ID",
    type="TRX_ID",
    value=1287,
    extern=nil>,
   #<struct Innodb::Page::Index::FieldDescriptor
    name="DB_ROLL_PTR",
    type="ROLL_PTR",
    value=
     #<struct Innodb::DataType::RollPointerType::Pointer
      is_insert=true,
      rseg_id=39,
      undo_log=#<struct Innodb::Page::Address page=283, offset=272>>,
    extern=nil>],
 child_page_number=nil,
 transaction_id=1287,
 roll_pointer=
  #<struct Innodb::DataType::RollPointerType::Pointer
   is_insert=true,
   rseg_id=39,
   undo_log=#<struct Innodb::Page::Address page=283, offset=272>>,
 length=49>


 % innodb_space -s /opt/homebrew/var/mysql/ibdata1  -T svc_tree/edge -p 3 page-illustrate
/System/Library/Frameworks/Ruby.framework/Versions/2.6/usr/lib/ruby/2.6.0/universal-darwin21/rbconfig.rb:230: warning: Insecure world writable dir /opt/homebrew in PATH, mode 040777

                                       Page 3 (INDEX)                          
      Offset ╭────────────────────────────────────────────────────────────────╮
           0 │█████████████████████████████████████▋██████████████████████████│
          64 │█████████▋███████████████████▋████████████▋████████████▋████▋███│
         128 │█████████████████████████████████████████████▋                  │
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
  █ Record Header                         5    0.03%
  █ Record Data                          49    0.30%
  █ Page Directory                        4    0.02%
  █ FIL Trailer                           8    0.05%
  ░ Garbage                               0    0.00%
    Free                              16198   98.86%

