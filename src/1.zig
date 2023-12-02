const std = @import("std");

pub fn main() !void {}

// const data_1a = .{ @embedFile("./input/1/test-a.txt"), 142 };
// const data_1b = .{ @embedFile("./input/1/test-b.txt"), 281 };
const data_1a = .{ @embedFile("./input/1/input.txt"), 55477 };
const data_1b = .{ @embedFile("./input/1/input.txt"), 54431 };

pub fn solve_1a(d: []const u8) i32 {
    var it = std.mem.split(u8, d, "\n");
    var sum: i32 = 0;
    while (it.next()) |x| {
        var first: ?u8 = null;
        var last: ?u8 = null;
        for (x) |v| {
            if ((v >= 48) and (v <= 57)) {
                if (first == null) {
                    first = v;
                }
                last = v;
            }
        }
        if (first != null and last != null) {
            var val = (first.? - '0') * @as(i32, 10) + (last.? - '0');
            // std.debug.print("{s} {d}\n", .{ x, val });
            sum += val;
        }
    }

    return sum;
}

pub fn solve_1b(d: []const u8) i32 {
    var it = std.mem.split(u8, d, "\n");
    var sum: i32 = 0;
    // WHY
    while (it.next()) |x| {
        var first: ?i32 = null;
        var last: ?i32 = null;
        for (x, 0..) |v, i| {
            if (((v >= 48) and (v <= 57))) {
                if (first == null) {
                    first = v - '0';
                }
                last = v - '0';
            } else if ((v == 'o') and (i + 2 < x.len) and (x[i + 1] == 'n') and (x[i + 2] == 'e')) {
                if (first == null) {
                    first = 1;
                }
                last = 1;
            } else if ((v == 't') and (i + 2 < x.len)) {
                if ((x[i + 1] == 'w') and (x[i + 2] == 'o')) {
                    if (first == null) {
                        first = 2;
                    }
                    last = 2;
                } else if ((i + 4 < x.len) and (x[i + 1] == 'h') and (x[i + 2] == 'r') and (x[i + 3] == 'e') and (x[i + 4] == 'e')) {
                    if (first == null) {
                        first = 3;
                    }
                    last = 3;
                }
            } else if ((v == 'f') and (i + 3 < x.len)) {
                if ((x[i + 1] == 'o') and (x[i + 2] == 'u') and (x[i + 3] == 'r')) {
                    if (first == null) {
                        first = 4;
                    }
                    last = 4;
                } else if ((x[i + 1] == 'i') and (x[i + 2] == 'v') and (x[i + 3] == 'e')) {
                    if (first == null) {
                        first = 5;
                    }
                    last = 5;
                }
            } else if ((v == 's') and (i + 2 < x.len)) {
                if ((x[i + 1] == 'i') and (x[i + 2] == 'x')) {
                    if (first == null) {
                        first = 6;
                    }
                    last = 6;
                } else if ((i + 4 < x.len) and (x[i + 1] == 'e') and (x[i + 2] == 'v') and (x[i + 3] == 'e') and (x[i + 4] == 'n')) {
                    if (first == null) {
                        first = 7;
                    }
                    last = 7;
                }
            } else if ((v == 'e') and (i + 4 < x.len) and (x[i + 1] == 'i') and (x[i + 2] == 'g') and (x[i + 3] == 'h') and (x[i + 4] == 't')) {
                if (first == null) {
                    first = 8;
                }
                last = 8;
            } else if ((v == 'n') and (i + 3 < x.len) and (x[i + 1] == 'i') and (x[i + 2] == 'n') and (x[i + 3] == 'e')) {
                if (first == null) {
                    first = 9;
                }
                last = 9;
            }
        }
        if (first != null and last != null) {
            var val = first.? * @as(i32, 10) + last.?;
            // std.debug.print("{s} {d}\n", .{ x, val });
            sum += val;
        }
    }

    return sum;
}

test "day1" {
    try std.testing.expectEqual(@as(i32, data_1a[1]), solve_1a(data_1a[0]));
    try std.testing.expectEqual(@as(i32, data_1b[1]), solve_1b(data_1b[0]));
}
