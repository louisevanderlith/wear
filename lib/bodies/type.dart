import 'size.dart';

class Type {
  final String name;
  final List<Size> sizes;

  Type(this.name, this.sizes);

  Map<String, dynamic> toJson() {
    return {
      "Name": name,
      "Sizes": sizes,
    };
  }
}
