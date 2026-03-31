open class Rectangle(x : Int, y : Int, hauteur : Int, largeur : Int) : FormeGeo(x,y,hauteur,largeur){
    override fun perimetre(): Double{
        return 2*dimension1.toDouble() + 2*dimension2.toDouble()
    }
    override fun surface() : Double{
        return dimension1*dimension2.toDouble()
    }
    override fun dessiner() : String{
        return "[${super.dessiner()},H=$dimension1,L=$dimension2]"
    }
}