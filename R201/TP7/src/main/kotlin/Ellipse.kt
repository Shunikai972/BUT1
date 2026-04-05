import kotlin.math.*

open class Ellipse(x : Int, y: Int, axeVertical: Int, axeHorizontal : Int) : FormeGeo(x,y,axeVertical, axeHorizontal){
    override fun perimetre(): Double {
        return 2*kotlin.math.PI*(kotlin.math.sqrt(((dimension1/2*dimension1/2).toDouble()+(dimension2/2*dimension2/2).toDouble())/2))
    }

    override fun surface(): Double {
        return kotlin.math.PI*(dimension1/2)*(dimension2/2)
    }

    override fun dessiner(): String {
        return "(${super.dessiner()},AV=$dimension1,AH=$dimension2)"
    }
}


// use kotlin.math.PI,  et kotlin.math.sqrt(...) si necessaire