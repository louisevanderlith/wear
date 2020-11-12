class Clothing {
  final String type;
  final String colour;
  final String size;

  Clothing(this.type, this.colour, this.size);

  Map<String, dynamic> toJson() {
    return {
      "Type": type,
      "Colour": colour,
      "Size": size,
    };
  }
}
