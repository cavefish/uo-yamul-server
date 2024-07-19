package dev.cavefish.yamul.backend.infra.localfile

import dev.cavefish.yamul.backend.game.controller.domain.Hue
import dev.cavefish.yamul.backend.game.controller.infra.mul.HueMulRepository
import org.springframework.stereotype.Service

@Service
class LocalFileHueMulRepository: HueMulRepository {
    override fun map(hue: Hue): Int {
        // https://uo.stratics.com/heptazane/fileformats.shtml#1.2

        return 0
    }
}