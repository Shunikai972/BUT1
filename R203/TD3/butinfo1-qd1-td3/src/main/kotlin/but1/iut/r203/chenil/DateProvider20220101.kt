package but1.iut.r203.chenil

import java.time.LocalDate

/**
 * Stub pour les cas CT_age2, CT_age3, CT_age4 (Exercice 4.2.1)
 * Retourne toujours la date 2022-01-01
 */
class DateProvider20220101 : DateProvider {
    override fun getDate(): LocalDate = LocalDate.of(2022, 1, 1)
}
