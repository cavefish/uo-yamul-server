package dev.cavefish.yamul

import com.flextrade.jfixture.JFixture
import org.assertj.core.api.SoftAssertions
import org.assertj.core.api.junit.jupiter.SoftAssertionsExtension
import org.junit.jupiter.api.AfterEach
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.extension.ExtendWith

@ExtendWith(SoftAssertionsExtension::class)
open class UnitTest {
    protected var softly: SoftAssertions = SoftAssertions()

    protected var fixture: JFixture = JFixture()

    @BeforeEach
    fun beforeEach() {
        fixture = JFixture()
    }

    @AfterEach
    fun afterEach() {
        softly.assertAll()
    }
}