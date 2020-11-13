class Clothing {
  final String code;
  final String brand;
  final String description;
  final String type;
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
