import 'dart:convert';
import 'dart:html';

import 'package:mango_ui/keys.dart';
import 'package:mango_ui/requester.dart';

import 'bodies/clothing.dart';

Future<HttpRequest> createClothing(Clothing obj) async {
  var apiroute = getEndpoint("wear");
  var url = "${apiroute}/info";

  return invokeService("POST", url, jsonEncode(obj.toJson()));
}

Future<HttpRequest> updateClothing(Key key, Clothing obj) async {
  var route = getEndpoint("wear");
  var url = "${route}/info/${key.toJson()}";

  final data = jsonEncode(obj.toJson());

  return invokeService("PUT", url, data);
}

Future<HttpRequest> deleteClothing(Key key) async {
  var route = getEndpoint("wear");
  var url = "${route}/info/${key.toJson()}";

  return invokeService("DELETE", url, "");
}
