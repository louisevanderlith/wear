import 'package:mango_ui/keys.dart';

class Clothing {
  final String code;
  final Key brand;
  final String description;
  final Key type;
  final String size;
  final String colour;
  final String material;
  final num weight;

  Clothing(this.code, this.brand, this.description, this.type, this.size,
      this.colour, this.material, this.weight);

  Map<String, dynamic> toJson() {
    return {
      "Code": code,
      "Brand": brand,
      "Description": description,
      "Type": type,
      "Size": size,
      "Colour": colour,
      "Material": material,
      "Weight": weight,
    };
  }
}
