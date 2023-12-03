const std = @import("std");

pub fn split(x: []const u8, c: []const u8) std.mem.SplitIterator(u8) {
    return std.mem.split(u8, x, c);
}

pub fn trim_char(x: []const u8, c: []const u8) []const u8 {
    return std.mem.trim(u8, x, c);
}

pub fn trim(x: []const u8) []const u8 {
    return std.mem.trim(u8, x, " ");
}

pub fn int(x: []const u8) !i32 {
    return std.fmt.parseInt(i32, x, 10);
}
