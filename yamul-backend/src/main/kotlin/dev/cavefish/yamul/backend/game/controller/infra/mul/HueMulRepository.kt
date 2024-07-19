package dev.cavefish.yamul.backend.game.controller.infra.mul

import dev.cavefish.yamul.backend.game.controller.domain.Hue

interface HueMulRepository {
    fun map(hue: Hue): Int
}