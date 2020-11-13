class Size {
  final String shortcode;
  final String description;
  final String dimensions;

  Size(this.shortcode, this.description, this.dimensions);

  Map<String, dynamic> toJson() {
    return {
      "ShortCode": shortcode,
      "Description": description,
      "Dimension": dimensions,
    };
  }
}
