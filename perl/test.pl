# -*- coding: utf-8 -*-
# perl

use strict;
use utf8; # necessary if you want to use Unicode in function or var names

# processing Unicode string
my $s = 'I ★ you';
$s =~ s/★/♥/;
print "$s\n";

# variable with Unicode char
my $愛 = 4;
print "$愛\n";

# function with Unicode char
sub f愛 { return 2;}
print f愛();