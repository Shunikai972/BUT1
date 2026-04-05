abstract class FormeGeo(x : Int, y : Int, dim1 : Int, dim2:  Int) : ObjetDessinable(x,y){
    protected var dimension1 : Int = dim1
    protected var dimension2 : Int = dim2
    abstract fun perimetre(): Double
    abstract fun surface(): Double
}