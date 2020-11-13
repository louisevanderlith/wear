class Brand {
  final String name;

  Brand(this.name);

  Map<String, dynamic> toJson() {
    return {
      "Name": name,
    };
  }
}
